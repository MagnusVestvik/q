use crate::color::Color;

pub struct Printer {
    folder_color: Color,
    files_color: Color,
    border_color: Color,
    show_size: bool,
}

impl Printer {
    fn new(
        folder_color: String,
        files_color: String,
        border_color: String,
        show_size: bool,
    ) -> Printer {
        Printer {
            folder_color: folder_color,
            files_color: files_color,
            border_color: border_color,
            show_size: show_size,
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
}
