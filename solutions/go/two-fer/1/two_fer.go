// package containing methods that show sharing is caring
package twofer

// ShareWith function returns a string describing the share procedure,
// usign the name provider, or a generic 'you' if no name is supplied
func ShareWith(name string) string {
	if name == "" {
        name = "you"
    }
    return "One for " + name + ", one for me."
}
