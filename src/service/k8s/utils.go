package k8s

var commonLabels = map[string]string{
	"github.com/migueleliasweb/d2k": "enabled",
}

//
func addCommonLabels(m map[string]string) map[string]string {
	for k, v := range commonLabels {
		m[k] = v
	}

	return m
}
