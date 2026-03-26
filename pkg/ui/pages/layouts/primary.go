package layouts

import (
	. "github.com/namzug16/gotags"
	"github.com/namzug16/suxless/pkg/ui"
	"github.com/namzug16/suxless/pkg/ui/basecoat"
	"github.com/namzug16/suxless/pkg/ui/lucide"
)

func Primary(title, currentPath string, content HTML) HTML {
	return Doc(
		Head(
			Metatags(title),
			CSS(),
			JS(),
		),
		Body(
			X.Class("min-h-screen bg-background"),
			Div(
				X.Class("flex min-h-screen"),
				sidebarMenu(currentPath),
				Main(
					X.Id("content"),
					X.Class("flex min-h-screen min-w-0 flex-1 flex-col"),
					Header(
						X.Class("bg-background sticky inset-x-0 top-0 isolate z-10 flex shrink-0 items-center gap-2 border-b"),
						Div(
							X.Class("flex h-14 w-full items-center gap-2 px-4"),
							Button(
								X.Type("button"),
								X.Attr("onclick", "document.dispatchEvent(new CustomEvent('basecoat:sidebar'))"),
								X.Attr("aria-label", "Toggle sidebar"),
								X.Attr("data-tooltip", "Toggle sidebar"),
								X.Attr("data-side", "bottom"),
								X.Attr("data-align", "start"),
								X.Class("btn-sm-icon-ghost size-7"),
								lucide.PanelLeft(),
							),
							H1(
								X.Class("text-sm font-medium"),
								title,
							),
							Div(
								X.Class("ml-auto flex items-center gap-2"),
								Button(
									X.Type("button"),
									X.Attr("aria-label", "Toggle dark mode"),
									X.Attr("data-tooltip", "Toggle dark mode"),
									X.Attr("data-side", "bottom"),
									X.Attr("onclick", "document.dispatchEvent(new CustomEvent('basecoat:theme'))"),
									X.Class("btn-icon-outline size-8"),
									Span(X.Class("hidden dark:block"), lucide.Sun()),
									Span(X.Class("block dark:hidden"), lucide.Moon()),
								),
								A(
									X.Href("https://github.com/namzug16/suxless"),
									X.Class("btn-icon size-8"),
									X.Target("_blank"),
									X.Rel("noopener noreferrer"),
									X.Attr("data-tooltip", "GitHub repository"),
									X.Attr("data-side", "bottom"),
									X.Attr("data-align", "end"),
									lucide.Github(),
								),
							),
						),
					),
					Div(
						X.Class("flex-1 p-4 md:p-6"),
						content,
					),
				),
			),
			basecoat.Toaster(basecoat.ToasterParams{ID: "toaster"}),
		),
	)
}

func sidebarMenu(currentPath string) HTML {
	return basecoat.Sidebar(basecoat.SidebarParams{
		ID:    "primary-sidebar",
		Label: "Main navigation",
		Header: Div(
			X.Class("flex items-center gap-2"),
			Img(
				X.Src(ui.StaticFile("logo.png")),
				X.Alt("Suxless"),
				X.Class("size-6"),
			),
			Span(X.Class("font-medium"), "Suxless"),
		),
		Menu: []basecoat.SidebarItem{
			{Label: "Home", URL: "/", Current: currentPath == "/"},
			{Label: "Kitchen Sink", URL: "/kitchen-sink", Current: currentPath == "/kitchen-sink"},
		},
		Footer: "UI playground",
	})
}

func JS() []HTML {
	return []HTML{
		Script(X.Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"), X.Defer()),
		Script(X.Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), X.Defer()),
		Script(X.Src(ui.StaticFile("js/basecoat.all.min.js")), X.Defer()),
		Script(
			Raw(`
    (() => {
      try {
        const stored = localStorage.getItem('themeMode');
        if (stored ? stored === 'dark'
                   : matchMedia('(prefers-color-scheme: dark)').matches) {
          document.documentElement.classList.add('dark');
        }
      } catch (_) {}

      const apply = dark => {
        document.documentElement.classList.toggle('dark', dark);
        try { localStorage.setItem('themeMode', dark ? 'dark' : 'light'); } catch (_) {}
      };

      document.addEventListener('basecoat:theme', (event) => {
        const mode = event.detail?.mode;
        apply(mode === 'dark' ? true
             : mode === 'light' ? false
             : !document.documentElement.classList.contains('dark'));
      });
    })();
		`),
		),
	}
}

func CSS() []HTML {
	return []HTML{
		Link(
			X.Href(ui.StaticFile("main.css")),
			X.Rel("stylesheet"),
			X.Type("text/css"),
		),
	}
}

func Metatags(title string) HTML {
	return Fragment(
		Meta(X.Charset("utf-8")),
		Meta(X.Name("viewport"), X.Content("width=device-width, initial-scale=1")),
		Link(X.Rel("icon"), X.Href(ui.StaticFile("favicon.png"))),
		Title(pageTitle(title)),
	)
}

func pageTitle(title string) string {
	if title == "" {
		return "Suxless"
	}

	return title + " | Suxless"
}
