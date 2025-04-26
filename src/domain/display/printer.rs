use super::color::Color;
use super::symbols::symbols as sym;
use string_builder::Builder;

pub struct Printer {
    folder_color: Color,
    files_color: Color,
    border_color: Color,
    show_size: bool,
}

// We should have name, type, size and modified. maybe skip the option of having size and make it
// default
//

/*
 * TODO: change string_builder in favour of write! macro since this is in std.
 * it can be used like the following:
 * write!(
 * &mut out,
 * "{}{}",
 * "smth1",
 * "smth2"
 * )
 * .unwrap();
 */

impl Printer {
    fn new(
        folder_color: Color,
        files_color: Color,
        border_color: Color,
        show_size: bool,
    ) -> Printer {
        Printer {
            folder_color,
            files_color,
            border_color,
            show_size,
        }
    }

    fn draw_input_vertically(input: Vec<String>) -> String {
        input.join("\n")
    }

    fn draw_input_show_size() -> String {
        sym::HORIZONTAL.to_string()
    }

    fn draw_header(&self, width: usize) -> String {
        let mut header: String = String::new();

        write!(
            &mut header,
            "{}{}{}\n",
            sym::TOP_LEFT,
            sym::HORIZONTAL.repeat(width - 2),
            sym::TOP_RIGHT,
        )
        .unwrap();

        write!(
            &mut header,
            "{}{}{}{}{}{}{}\n",
            sym::VERTICAL,
            "#", // TODO: this will be ugly looking should make somthing for boxing theese in
            "Name",
            "Type",
            "Size",
            "modified",
            sym::VERTICAL,
        )
        .unwrap();

        header
    }

    //  TODO: Implement
    fn draw_box(item: String, x_padding: usize, y_padding: usize, has_top: bool) -> String {
        let WIDTH: usize = x_padding * 2 + item.len();
        let mut box: String = String::new(); 

        "".to_string()
    }

    fn draw_bottom(width: usize) -> String {
        let mut sb = Builder::default();
        sb.append(format!(
            "{}{}{}",
            sym::BOTTOM_LEFT,
            sym::HORIZONTAL.repeat(width.saturating_sub(2)),
            sym::BOTTOM_RIGHT,
        ));

        sb.string().unwrap()
    }

    fn draw_body<F>(&self, input: Vec<Vec<String>>, mut draw_row: F) -> String
    where
        F: FnMut(Vec<String>, usize) -> String,
    {
        let mut to_draw: Vec<String> = input
            .iter()
            .enumerate()
            .map(|(i, row)| draw_row(row.to_vec(), i))
            .collect();

        to_draw.push(Self::draw_bottom(to_draw.len()));

        to_draw.join("\n")
    }
}
