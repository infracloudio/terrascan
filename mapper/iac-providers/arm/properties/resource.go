package properties

// Resource exposes common properties of a resource.
var Resource *resource

type resource struct {
	Location string
}

func init() {
	Resource = &resource{
		Location: "location",
	}
}
