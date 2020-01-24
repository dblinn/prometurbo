package provider

import (
	"strings"

	"github.com/golang/glog"
	"github.com/turbonomic/prometurbo/prometurbo/appmetric/internal/prometheus"
	"github.com/turbonomic/prometurbo/prometurbo/appmetric/metrics"
)

type Provider struct {
	promClients  []*prometheus.RestClient
	exporterDefs []*exporterDef
}

func NewProvider(promClients []*prometheus.RestClient, exporterDefs []*exporterDef) *Provider {
	return &Provider{
		promClients:  promClients,
		exporterDefs: exporterDefs,
	}
}

func (p *Provider) GetMetrics() ([]*metrics.EntityMetric, error) {
	var allMetrics []*metrics.EntityMetric
	// TODO: use goroutine
	for _, promClient := range p.promClients {
		var metricsForProms []*metrics.EntityMetric
		for _, exporterDef := range p.exporterDefs {
			metricsForExporters := getMetricsForExporter(promClient, exporterDef)
			metricsForProms = append(metricsForProms, metricsForExporters...)
		}
		allMetrics = append(allMetrics, metricsForProms...)
	}
	return allMetrics, nil
}

func (p *Provider) Validate() error {
	for _, promClient := range p.promClients {
		if _, err := promClient.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Provider) String() string {
	hosts := make([]string, len(p.promClients))
	for i, client := range p.promClients {
		hosts[i] = client.GetHost()
	}
	return strings.Join(hosts, ", ")
}

func getMetricsForExporter(
	promClient *prometheus.RestClient, exporterDef *exporterDef) []*metrics.EntityMetric {
	var metricsForExporter []*metrics.EntityMetric
	for _, entityDef := range exporterDef.entityDefs {
		metricsForEntity := getMetricsForEntity(promClient, entityDef)
		metricsForExporter = append(metricsForExporter, metricsForEntity...)
	}
	return metricsForExporter
}

func getMetricsForEntity(
	promClient *prometheus.RestClient, entityDef *entityDef) []*metrics.EntityMetric {
	var metricsForEntity []*metrics.EntityMetric
	var metricsForEntityMap = map[string]*metrics.EntityMetric{}
	for metricType, metricDef := range entityDef.metricDefs {
		for metricKind, metricQuery := range metricDef.queries {
			metricSeries, err := promClient.GetMetrics(metricQuery)
			if err != nil {
				glog.Errorf("Failed to query metric %v [%v] for entity type %v: %v.",
					metricKind, metricQuery, entityDef.eType, err)
				continue
			}
			for _, metricData := range metricSeries {
				basicMetricData, ok := metricData.(*prometheus.BasicMetricData)
				if !ok {
					// TODO: Enhance error messages
					glog.Errorf("Type assertion failed for metricData %+v obtained from %v [%v] for entity type %v.",
						metricData, metricKind, metricQuery, entityDef.eType)
					continue
				}
				id, attr, err := entityDef.reconcileAttributes(basicMetricData.Labels)
				if err != nil {
					glog.Errorf("Failed to reconcile attributes from labels %+v obtained from %v [%v] for entity %v: %v.",
						basicMetricData.Labels, metricKind, metricQuery, entityDef.eType, err)
					continue
				}
				if id == "" {
					glog.Warningf("Failed to get identifier from labels %+v obtained from %v [%v] for entity %v.",
						basicMetricData.Labels, metricKind, metricQuery, entityDef.eType)
					continue
				}

				if _, found := metricsForEntityMap[id]; !found {
					metricsForEntityMap[id] = metrics.NewEntityMetric(id, entityDef.eType)
				}
				for name, value := range attr {
					metricsForEntityMap[id].SetLabel(name, value)
				}
				metricsForEntityMap[id].SetMetric(metricType, metricKind, basicMetricData.GetValue())
			}
		}
	}
	for _, metric := range metricsForEntityMap {
		metricsForEntity = append(metricsForEntity, metric)
	}
	return metricsForEntity
}
