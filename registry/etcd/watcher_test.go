package etcd

import (
	"os"
	"testing"
	"time"

	"github.com/hb-chen/gateway/registry"
)

func TestWatcher(t *testing.T) {
	if travis := os.Getenv("TRAVIS"); travis == "true" {
		t.Skip()
	}

	testData := []*registry.Service{
		{
			Name:    "test1",
			Version: "1.0.1",
			Methods: []*registry.Method{
				{
					Name: "test1-1",
					Routes: []*registry.Route{
						{
							Method: "GET",
							Pattern: &registry.Pattern{
								Version:         1,
								Ops:             nil,
								Pool:            nil,
								Verb:            "",
								AssumeColonVerb: false,
							},
						},
					},
				},
			},
		},
		{
			Name:    "test1",
			Version: "1.0.2",
			Methods: []*registry.Method{
				{
					Name: "test1-2",
					Routes: []*registry.Route{
						{
							Method: "POST",
							Pattern: &registry.Pattern{
								Version:         1,
								Ops:             nil,
								Pool:            nil,
								Verb:            "",
								AssumeColonVerb: false,
							},
						},
					},
				},
			},
		},
		{
			Name:    "test2",
			Version: "1.0.1",
			Methods: []*registry.Method{
				{
					Name: "test2-1",
					Routes: []*registry.Route{
						{
							Method: "GET",
							Pattern: &registry.Pattern{
								Version:         1,
								Ops:             nil,
								Pool:            nil,
								Verb:            "",
								AssumeColonVerb: false,
							},
						},
					},
				},
			},
		},
	}

	testFn := func(service, s *registry.Service) {
		if s == nil {
			t.Fatalf("Expected one result for %s got nil", service.Name)

		}

		if s.Name != service.Name {
			t.Fatalf("Expected name %s got %s", service.Name, s.Name)
		}

		if s.Version != service.Version {
			t.Fatalf("Expected version %s got %s", service.Version, s.Version)
		}

		if len(s.Methods) != 1 {
			t.Fatalf("Expected 1 node, got %d", len(s.Methods))
		}

		method := s.Methods[0]

		if method.Name != service.Methods[0].Name {
			t.Fatalf("Expected node id %s got %s", service.Methods[0].Name, method.Name)
		}

		if method.Routes[0].Method != service.Methods[0].Routes[0].Method {
			t.Fatalf("Expected node address %s got %s", service.Methods[0].Routes[0].Method, method.Routes[0].Method)
		}
	}

	travis := os.Getenv("TRAVIS")

	var opts []registry.Option

	if travis == "true" {
		opts = append(opts, registry.Timeout(time.Millisecond*100))
	}

	// new registry
	r := NewRegistry()

	w, err := r.Watch()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Stop()

	for _, service := range testData {
		// register service
		if err := r.Register(service); err != nil {
			t.Fatal(err)
		}

		for {
			res, err := w.Next()
			if err != nil {
				t.Fatal(err)
			}

			if res.Service.Name != service.Name {
				continue
			}

			if res.Action != "create" {
				t.Fatalf("Expected create event got %s for %s", res.Action, res.Service.Name)
			}

			testFn(service, res.Service)
			break
		}

		// deregister
		if err := r.Deregister(service); err != nil {
			t.Fatal(err)
		}

		for {
			res, err := w.Next()
			if err != nil {
				t.Fatal(err)
			}

			if res.Service.Name != service.Name {
				continue
			}

			if res.Action != "delete" {
				continue
			}

			testFn(service, res.Service)
			break
		}
	}
}
