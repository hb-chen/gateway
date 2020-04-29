package registry

// CopyService make a copy of service
func CopyService(service *Service) *Service {
	// copy service
	s := new(Service)
	*s = *service

	// copy nodes
	methods := make([]*Method, len(service.Methods))
	for i, method := range service.Methods {
		m := new(Method)
		*m = *method

		// copy routes
		routes := make([]*Route, len(method.Routes))
		for j, route := range method.Routes {
			r := new(Route)
			*r = *route

			p := new(Pattern)
			*p = *route.Pattern
			r.Pattern = p

			routes[j] = r
		}
		m.Routes = routes

		methods[i] = m
	}
	s.Methods = methods

	return s
}

// Copy makes a copy of services
func Copy(current []*Service) []*Service {
	services := make([]*Service, len(current))
	for i, service := range current {
		services[i] = CopyService(service)
	}
	return services
}
