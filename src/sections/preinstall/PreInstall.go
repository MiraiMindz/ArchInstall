package preinstall

import (
	"fmt"
	"os"

	"golang.org/x/term"

	dms "utils/ui/components/Descriptive/MultiSelect"
	ds "utils/ui/components/Descriptive/Select"
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
		true,
		lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Height(h).Width(w).Padding(2),
		"# Markdown Rendered Important Message",
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

func sectionF(w, h int) string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Select Option  [ENTER]: Send Response")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := ds.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		ds.Select(lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top), (w/2)-4, h-(4+hps+hds), "Select Something",
			ds.Option("A", "Test Description 00"),
			ds.Option("B", "Test Description 01"),
			ds.Option("C", "Test Description 02\nThis is a multiline\nDescription to resize items"),
			ds.Option("D", "Test Description 03"),
			ds.Option("E", "Test Description 04"),
			ds.Option("F", "Test Description 05"),
			ds.Option("G", "Test Description 06"),
			ds.Option("H", "Test Description 07"),
			ds.Option("I", "Test Description 08"),
			ds.Option("J", "Test Description 09"),
			ds.Option("K", "Test Description 10"),
			ds.Option("L", "Test Description 11"),
			ds.Option("M", "Test Description 12"),
			ds.Option("N", "Test Description 13"),
			ds.Option("O", "Test Description 14"),
			ds.Option("P", "Test Description 15"),
			ds.Option("Q", "Test Description 16"),
			ds.Option("R", "Test Description 17"),
			ds.Option("S", "Test Description 18"),
			ds.Option("T", "Test Description 19"),
			ds.Option("U", "Test Description 20"),
			ds.Option("V", "Test Description 21"),
			ds.Option("W", "Test Description 22"),
			ds.Option("X", "Test Description 23"),
			ds.Option("Y", "Test Description 24"),
			ds.Option("Z", "Test Description 25"),
		),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide B", sideContent),
		hdt,
		hptl,
		hptr,
	))))

	return r
}

func sectionG(w, h int) []string {
	hptl, hps := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Cursor Up    [DOWN]: Move Cursor Down\n[SPACE]: Select Option  [ENTER]: Send Response")
	hptr, _ := helpmenu.HelpMenu("[TAB]: Switch Focus\n[UP]: Move Page Up    [DOWN]: Move Page Down\n")
	hdt, hds := header.Header("[03/26]\tPre-Install\tSelect Window Manager")

	r := dms.GetAnswer(spl.GetLeftSide(exc.Executor(spl.SplitScreen(
		dms.MultiSelect(lipgloss.NewStyle().Height(h-(4+hps+hds)).Width((w/2)-4).Align(lipgloss.Left, lipgloss.Top), (w/2)-4, h-(4+hps+hds), "Select Something",
			dms.Option("A", "Test Description 00"),
			dms.Option("B", "Test Description 01"),
			dms.Option("C", "Test Description 02\nThis is a multiline\nDescription to resize items"),
			dms.Option("D", "Test Description 03"),
			dms.Option("E", "Test Description 04"),
			dms.Option("F", "Test Description 05"),
			dms.Option("G", "Test Description 06"),
			dms.Option("H", "Test Description 07"),
			dms.Option("I", "Test Description 08"),
			dms.Option("J", "Test Description 09"),
			dms.Option("K", "Test Description 10"),
			dms.Option("L", "Test Description 11"),
			dms.Option("M", "Test Description 12"),
			dms.Option("N", "Test Description 13"),
			dms.Option("O", "Test Description 14"),
			dms.Option("P", "Test Description 15"),
			dms.Option("Q", "Test Description 16"),
			dms.Option("R", "Test Description 17"),
			dms.Option("S", "Test Description 18"),
			dms.Option("T", "Test Description 19"),
			dms.Option("U", "Test Description 20"),
			dms.Option("V", "Test Description 21"),
			dms.Option("W", "Test Description 22"),
			dms.Option("X", "Test Description 23"),
			dms.Option("Y", "Test Description 24"),
			dms.Option("Z", "Test Description 25"),
		),
		pg.Pager(true, w/2-4, h-(4+hps+hds), "Guide B", sideContent),
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
	f := sectionF(w, h)
	g := sectionG(w, h)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
