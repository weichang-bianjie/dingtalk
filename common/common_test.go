package common

import "testing"

func TestFilterContent(t *testing.T) {
	content := "{\"receiver\":\"web\\\\.hook\",\"status\":\"firing\",\"alerts\":[{\"status\":\"firing\",\"labels\":{\"alertname\":\"监控告警\",\"exported_job\":\"hwctest\",\"instance\":\"pushgateway\",\"job\":\"pushgateway\",\"name\":\"hwc\",\"sex\":\"man\"},\"annotations\":{\"description\":\"pushgateway of job pushgateway has been down for more than 5 minutes.\",\"summary\":\"Instance pushgateway down\"},\"startsAt\":\"2021-02-23T10:02:01.431338009Z\",\"endsAt\":\"0001-01-01T00:00:00Z\",\"generatorURL\":\"http://16cf2a66ea3c:9090/graph?g0.expr=some_metric+%3D%3D+0\\u0026g0.tab=1\",\"fingerprint\":\"ef48af42e13b1a95\"}],\"groupLabels\":{\"alertname\":\"监控告警\"},\"commonLabels\":{\"alertname\":\"监控告警\",\"exported_job\":\"hwctest\",\"instance\":\"pushgateway\",\"job\":\"pushgateway\",\"name\":\"hwc\",\"sex\":\"man\"},\"commonAnnotations\":{\"description\":\"pushgateway of job pushgateway has been down for more than 5 minutes.\",\"summary\":\"Instance pushgateway down\"},\"externalURL\":\"http://5f1e3f495741:9093\",\"version\":\"4\",\"groupKey\":\"{}:{alertname=\\\"监控告警\\\"}\",\"truncatedAlerts\":0}\n"
	t.Log(FilterContent([]byte(content)))
}
