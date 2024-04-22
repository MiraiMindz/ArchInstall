package preinstall

import (
	"fmt"
	"os"

	"golang.org/x/term"

	header "utils/ui/components/Header"
	helpmenu "utils/ui/components/HelpMenu"
	ms "utils/ui/components/MultiSelect"
	pg "utils/ui/components/Pager"
	sel "utils/ui/components/Select"
	spl "utils/ui/components/SplitScreen"
	te "utils/ui/components/TextElement"
	ti "utils/ui/components/TextInput"
	exc "utils/ui/meta/Executor"

	"github.com/charmbracelet/lipgloss"
)

const sideContent = `
# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] Rene Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appetit!

# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] Rene Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appetit!

# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] Rene Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appetit!

# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] Rene Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appetit!
`

func sectionA(w, h int) []string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Mark Option  [ENTER]: Send Response")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := ms.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		ms.MultiSelect(lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top), (w/2)-4, h-(4+hps+hds), "MULTI Select Something", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide A", sideContent),
		hdt,
		hptl,
		hptr,
	))))

	return r
}

func sectionB(w, h int) string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Select Option  [ENTER]: Send Response")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := sel.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		sel.Select(lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top), (w/2)-4, h-(4+hps+hds), "Select Something", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide B", sideContent),
		hdt,
		hptl,
		hptr,
	))))

	return r
}

func sectionC(w, h int) string {
	r := te.GetAnswer(exc.Executor(te.TextElement(
		lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(h).Width(w),
		"Important message",
	)))
	return r
}

func sectionD(w, h int) string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[ENTER]: Send Response\n")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := ti.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		ti.TextInput("Type Something:", "Insert here", false, 512, w/2-8, lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top)),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide D", sideContent),
		hdt,
		hptl,
		hptr,
	))))

	return r

}

func sectionE(w, h int) string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[ENTER]: Send Response\n")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := ti.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		ti.TextInput("Type Something Secret:", "Insert here", true, 512, w/2-8, lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top)),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide E", sideContent),
		hdt,
		hptl,
		hptr,
	))))

	return r

}

func PreInstall() {
	w, h, er := term.GetSize(int(os.Stdin.Fd()))
	if er != nil {
		panic(er)
	}

	a := sectionA(w, h)
	b := sectionB(w, h)
	c := sectionC(w, h)
	d := sectionD(w, h)
	e := sectionE(w, h)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
