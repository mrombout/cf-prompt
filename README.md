# cf-prompt

A utility to create custom shell prompts for Cloud Foundry.

## Usage

Use `cf-prompt` inside the value of your `PROMPT` environment variable:

```
$ PROMPT='$(cf-prompt) $'
Organization / Space $ 
```

Use the `-t` flag to set a custom template. Refer to [text/template](https://golang.org/pkg/text/template/) for more information about the templating language.

```
  -t string
        The text/template to use for output.
    
        Available values are:
    
                - {{.AccessToken}}
                - {{.OrganizationFields.Name}}
                - {{.OrganizationFields.URL}}
                - {{.SpaceFields.Name}}
         (default "{{if .AccessToken}}{{.OrganizationFields.Name}}{{if .SpaceFields.Name}} / {{.SpaceFields.Name}}{{end}}{{end}}")
  -x    Print errors and other output to aid in debugging
```
