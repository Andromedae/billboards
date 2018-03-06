package display


import (
	"strings"
)


func BestFontSize(width, height int , text string) (int, int) {

	//algorithm: start with a max font size (character fills all the height of the billboard. 
	//we assume billboard cannot be turned by 90°

	font_size := height
    nb_lines := 1

    if len(text) == 0 {
        return 0, 0
    }

    for  {
    	if fits(text, font_size, nb_lines, width) {
    		return font_size, nb_lines
    	} else { 
        	font_size, nb_lines = getNextFontSize(font_size, height, width, nb_lines)
    	}
    }

	return font_size, nb_lines
}


//helper function, returns true if the given string fits in the billboard, using the given size
func fits(text string, font_size, free_lines, width int) bool {

	letters := len(text)
    characters_per_line := width / font_size //nb of chars we can squeeze in one line

    if free_lines == 0 {
        return false
    }

    if letters <= characters_per_line {
        return true
    }

    last_space_index := strings.LastIndex(text, " ")
    return fits(text[last_space_index+1:], font_size, free_lines - 1, width)
}

//helper function, returns true if the given string fits in the billboard, using the given size
func fits2(text string, font_size, free_lines, width int) bool {

	words := strings.Split(text, " ")

	letters := len(text)
    characters_per_line := width / font_size //nb of chars we can squeeze in one line

    if free_lines == 0 {
        return false
    }

    if letters <= characters_per_line {
        return true
    }

    last_space_index := strings.LastIndex(text, " ")
    return fits(text[last_space_index+1:], font_size, free_lines - 1, width)
}



func getNextFontSize(font_size, height, width , nb_lines int ) (int, int) {
    characters_per_line := width / font_size
    new_font_size_on_width := width / (characters_per_line + 1)

    new_font_size_on_height := height / (nb_lines + 1)

    if (new_font_size_on_width > new_font_size_on_height) {
        return new_font_size_on_width, nb_lines
    } else {
        return new_font_size_on_height, nb_lines + 1
    }
}

