package sybil

var VERSION_STRING = "0.4.0"

func GetVersionInfo() map[string]interface{} {
	version_info := make(map[string]interface{})

	version_info["version"] = VERSION_STRING
	version_info["field_separator"] = true
	version_info["log_hist"] = true
	version_info["query_cache"] = true

	return version_info

}
