package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type ExtraLink struct {
	Title string `yaml:"Title"`
	Url   string `yaml:"URL"`
}

type Link struct {
	Title      string      `yaml:"Title"`
	Url        string      `yaml:"URL"`
	ExtraLinks []ExtraLink `yaml:"ExtraLinks"`
}

type Block struct {
	Heading string `yaml:"Heading"`
	Links   []Link `yaml:"Links"`
}

type Col struct {
	Blocks []Block `yaml:"Blocks"`
}

type Row struct {
	Cols []Col `yaml:"Cols"`
}

type Config struct {
	SiteHeading      string        `yaml:"SiteHeading"`
	SiteTitle        string        `yaml:"SiteTitle"`
	CompanyName      string        `yaml:"CompanyName"`
	CompanyDomain    string        `yaml:"CompanyDomain"`
	CompanyUrl       string        `yaml:"CompanyUrl"`
	Rows             []Row         `yaml:"Rows"`
	ExtraFooterLinks []Link        `yaml:"ExtraFooterLinks"`
	HtmlHeadExtra    template.HTML `json:"HtmlHeadExtra" yaml:"HtmlHeadExtra"`
}

var TEMPLATE = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.SiteTitle}}</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
    {{.HtmlHeadExtra}}
    <style>
      body {
        padding: 2em 2em;
        margin: 0;
        font-size: 1.3em;
        font-weight: 400;
        line-height: 1.5;
        background-color: #fff;
        -webkit-text-size-adjust: 100%;
        -webkit-tap-highlight-color: transparent;
        font-family: "IBM Plex Sans", sans-serif;
      }
      h1 {
        margin: 0;
      }
      p {
        margin: 0em 1em 1em 0;
      }
      a.button{
        border: 2px solid #131480;
        padding: 0.15em 0.4em;
        color: #131480;
        -webkit-text-decoration: none;
        text-decoration: none;
        display: inline-block;
      }
      a.button:hover{
        text-decoration: none;
        border: 2px solid #131480;
        background-color: #131480;
        color:white;
      }
      .row{
        --bs-gutter-y: 0;
        display: -moz-box;
        display: flex;
        flex-wrap: wrap;
        margin-top: calc(var(--bs-gutter-y) * -1);
        margin-right: calc(var(--bs-gutter-x) * -.5);
        margin-left: calc(var(--bs-gutter-x) * -.5);
      }
      .row>* {
        width: 100%;
        max-width: 100%;
        padding-right: calc(var(--bs-gutter-x) * .5);
        padding-left: calc(var(--bs-gutter-x) * .5);
        margin-top: var(--bs-gutter-y);
      }
      @media (min-width: 768px){
        .column {
          width: 33.33333333%;
          flex: 0 0 auto;
        }
      }
      .footer {
        margin-top: 2em;
      }
      .footer p {
        margin: 0;
        }
      .right {
        text-align: right;
      }
    </style>
   </head>
   <body>
      <h1>{{.SiteHeading}}</h1>
      {{range $_, $row := .Rows}}
        <div class="row">
          {{range $_, $col := $row.Cols}}
            <div class="column">
                {{range $_, $block := $col.Blocks}}
                  <h2>{{.Heading}}</h2>
                  {{range $_, $link := $block.Links}}
                    <p>
											<a class="button" target="_blank" rel="noopener noreferrer" href="{{$link.Url}}" >{{$link.Title}}</a>
											{{range $_, $extraLink := $link.ExtraLinks}}
												<a class="button" target="_blank" rel="noopener noreferrer" href="{{$extraLink.Url}}" >{{$extraLink.Title}}</a>
											{{end}}
										</p>
                  {{end}}
                {{end}}
            </div>
          {{end}}
        </div>
      {{end}}
      </div>
      <div class="footer">
         <p>
            <strong>
               {{.CompanyName}}{{ if and (.CompanyDomain) (.CompanyUrl) }}, <a href="{{.CompanyUrl}}" target="_blank" style="color:black">{{.CompanyDomain}}</a>{{end}}
            </strong>
         </p>
         {{ if .ExtraFooterLinks }}
				 <p style="font-size: 0.8em">
            {{ range $_, $link := .ExtraFooterLinks }}
              <a href="{{$link.Url}}" target="_blank" style="color:black">{{$link.Title}}</a>
            {{ end }}
         </p>
         {{ end }}
         <p class="right">
            <a href="https://github.com/sikalabs/signpost" target="_blank" style="color:black">signpost</a> by <a href="https://sikalabs.com" target="_blank" style="color:black">sikalabs</a>
         </p>
      </div>
   </body>
</html>
`

var HTML string

func Server(config Config) error {
	if config.SiteTitle == "" {
		config.SiteTitle = config.SiteHeading
	}

	t := template.Must(template.New("index-html").Parse(TEMPLATE))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, config)
	if err != nil {
		return err
	}
	HTML = tpl.String()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, HTML)
	})
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK\n")
	})
	http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK\n")
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := 8000
	fmt.Printf("Server started on 0.0.0.0:%d, see http://127.0.0.1:%d\n", port, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
