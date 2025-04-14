use super::color::Color;
use super::symbols;
use string_builder::Builder;

pub struct Printer {
    folder_color: Color,
    files_color: Color,
    border_color: Color,
    show_size: bool,
}

// We should have name, type, size and modified. maybe skip the option of having size and make it
// default

impl Printer {
    // TODO: missing logic for actually drawing header text
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

    fn draw(&self, path: String, get_files_and_folders: impl Fn(String) -> Vec<String>) -> String {
        let to_render: Vec<String> = get_files_and_folders(path);
        if self.show_size {
            // TODO: implement
        }
        "".to_string()
    }

    fn draw_input_vertically(input: Vec<String>) -> String {
        input.join("\n")
    }

    fn draw_input_show_size() -> String {
        symbols::symbols::HORIZONTAL.to_string()
    }

    fn draw_header(&self, width: usize, height: usize) -> String {
        let mut sb = Builder::default();
        sb.append(format!(
            "{}{}{}\n",
            symbols::symbols::TOP_LEFT,
            symbols::symbols::HORIZONTAL.repeat(width - 2),
            symbols::symbols::TOP_RIGHT,
        ));
        sb.append(format!(
            // TODO: should create boxes for each column here
            "{}{}{}{}{}{}{}",
            symbols::symbols::VERTICAL,
            "#",
            "Name",
            "Type",
            "Size",
            "modified",
            symbols::symbols::VERTICAL,
        ));

        sb.string().unwrap()
    }

    fn draw_body(&self, input: Vec<String>) -> String {
        "".to_string()
    }
    fn draw_body_with_size(&self, input: Vec<String>, size: Vec<String>) -> String {
        "".to_string()
    }

    fn draw_footer(&self) -> String {
        "".to_string()
    }

    fn draw_box() -> String {
        let mut sb = Builder::default();
        "".to_string()
    }
}
