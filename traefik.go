package main

func ListRules() (map[string]Route, error) {
	var rules = map[string]Route{}

	services, err := ListServices()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		for _, route := range SplitMatchPath(service.Rule) {
			key := route.ToString()
			if key != "" {
				rules[key] = route
			}
		}
	}

	return rules, nil

}
