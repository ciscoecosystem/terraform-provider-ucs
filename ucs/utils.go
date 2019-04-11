package ucs

import "strings"

func GetParentDn(childDn string) string {
	arr := strings.Split(childDn, "/")
	// in case of cidr blocks we have extra / in the ip range so let's catch it and remove. This will have extra part.
	lastElement := arr[len(arr)-1]
	if strings.Contains(lastElement, "]") {
		backSlashedDn := strings.Join(arr[:len(arr)-1], "/")
		// split on - to remove last element.
		arr = strings.Split(backSlashedDn, "/")
		// remove last 2 elements as that will contain the extra part like rn - ip

		dnWithSlash := strings.Join(arr[:len(arr)-1], "/")

		return strings.TrimSuffix(dnWithSlash, "/")

	}

	return strings.Join(arr[:len(arr)-1], "/")

}
