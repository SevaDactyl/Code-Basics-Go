/*
На веб-сайтах часто используются разные поддомены для языков. Например, сайт site.com на английском располагается по адресу en.site.com, а на русском — ru.site.com.

Реализуйте функцию DomainForLocale(domain, locale string) string, которая добавляет язык locale как поддомен к домену domain. Язык может прийти пустым, тогда нужно добавить поддомен en.. Например:

DomainForLocale("site.com", "") // en.site.com
DomainForLocale("site.com", "ru") // ru.site.com
*/

// Первое решение

package solution

import (
	"fmt"
)

// BEGIN (write your solution here)
func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}
	return fmt.Sprintf(locale + "." + domain)
} 

// END

// Второе решение

package solution

import (
	"fmt"
)

// BEGIN (write your solution here)

// DomainForLocale adds a subdomain to a domain depending on locale.
func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en"
	}

	return fmt.Sprintf("%s.%s", locale, domain)
}

// END
