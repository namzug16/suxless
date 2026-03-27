package pages

import (
	. "github.com/namzug16/gotags"
	"github.com/namzug16/suxless/pkg/ui/pages/layouts"
)

func Home() HTML {
	return layouts.Primary("Home", "/", homeContent())
}

func homeContent() HTML {
	return Div(
		X.Class("w-full space-y-8"),
		Div(
			X.Class("space-y-2"),
			H1(X.Class("text-3xl font-semibold tracking-tight"), "Suxless server is up"),
			P(
				X.Class("text-muted-foreground"),
				"Use this workspace to build pages and test basecoat components.",
			),
		),
		Div(
			X.Class("grid gap-4 md:grid-cols-2"),
			A(
				X.Href("/kitchen-sink"),
				X.Class("rounded-lg border p-5 no-underline hover:bg-muted/50 transition-colors block"),
				H2(X.Class("text-lg font-medium mb-1"), "Kitchen Sink"),
				P(X.Class("text-sm text-muted-foreground"), "Browse all available Basecoat components."),
			),
			Div(
				X.Class("rounded-lg border p-5"),
				H2(X.Class("text-lg font-medium mb-1"), "Health"),
				P(X.Class("text-sm text-muted-foreground"), "API health endpoint: /health"),
			),
		),
	)
}
