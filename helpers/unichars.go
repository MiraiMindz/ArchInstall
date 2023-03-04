package helpers

import (
	"strconv"
	"fmt"
)


func BoxHorizontalChar() string  {
	box_draw_h_char_code, _, _, err := strconv.UnquoteChar("\u2500", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_h_char_code)
}

func BoxHorizontalUpChar() string  {
	box_draw_h_up_char_code, _, _, err := strconv.UnquoteChar("\u2534", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_h_up_char_code)
}

func BoxVerticalChar() string  {
	box_draw_v_char_code, _, _, err := strconv.UnquoteChar("\u2502", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_v_char_code)
}

func BoxLeftBottomChar() string  {
	box_draw_lb_char_code, _, _, err := strconv.UnquoteChar("\u2514", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_lb_char_code)
}

func BoxLeftTopChar() string {
	box_draw_lt_char_code, _, _, err := strconv.UnquoteChar("\u250C", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_lt_char_code)
}

func BoxRightBottomChar() string  {
	box_draw_rb_char, _, _, err := strconv.UnquoteChar("\u2518", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_rb_char)
}

func BoxRightTopChar() string {
	box_draw_rt_char, _, _, err := strconv.UnquoteChar("\u2510", 0)
	Check(err)
	return fmt.Sprintf("%c", box_draw_rt_char)
}
