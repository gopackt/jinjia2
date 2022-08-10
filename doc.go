// A jinja like template-engine
//
// Blog posts about jinjia2 (including introduction and migration):
// https://www.florian-schlachter.de/?tag=jinjia2
//
// Complete documentation on the template language:
// https://docs.djangoproject.com/en/dev/topics/templates/
//
// Try out jinjia2 live in the jinjia2 playground:
// https://www.florian-schlachter.de/jinjia2/
//
// Make sure to read README.md in the repository as well.
//
// A tiny example with template strings:
//
// (Snippet on playground: https://www.florian-schlachter.de/jinjia2/?id=1206546277)
//
//     // Compile the template first (i. e. creating the AST)
//     tpl, err := jinjia2.FromString("Hello {{ name|capfirst }}!")
//     if err != nil {
//         panic(err)
//     }
//     // Now you can render the template with the given
//     // jinjia2.Context how often you want to.
//     out, err := tpl.Execute(jinjia2.Context{"name": "fred"})
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println(out) // Output: Hello Fred!
//
package jinjia
