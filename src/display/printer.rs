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

    // TODO: rename
    fn createDrawItems(&self, path: String) -> Vec<Vec<String>> {
        let names: Vec<String> = get_files_and_folders(path);
        let types: Vec<String> = get_file_types(&names);
        let sizes: Vec<String> = get_file_sizes(&names);

        names
            .iter()
            .zip(types.iter())
            .zip(sizes.iter())
            .map(|((name, ty), size)| vec![name.clone(), ty.clone(), size.clone()])
            .collect()
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

    // TODO: not extensive but works for now, maybe should take in a
    // draw function that would be responsible for what to draw on each line or somthing similar
    fn draw_body(&self, input: Vec<Vec<String>>) -> String {
        let mut sb = Builder::default();
        for i in 0..input.len() {
            sb.append(format!(
                "{}{}{}{}{}{}{}{}{}{}{}",
                symbols::symbols::VERTICAL,
                i,
                symbols::symbols::VERTICAL,
                input[i][0],
                symbols::symbols::VERTICAL,
                input[i][1],
                symbols::symbols::VERTICAL,
                input[i][2],
                symbols::symbols::VERTICAL,
                input[i][4],
                symbols::symbols::VERTICAL,
            ));
            if i + 1 == input.len() {
                sb.append(format!(
                    "{}{}{}",
                    symbols::symbols::BOTTOM_LEFT,
                    symbols::symbols::HORIZONTAL.repeat(input.len() - 2),
                    symbols::symbols::BOTTOM_RIGHT,
                ));
            }
        }
        sb.string().unwrap()
    }
}

fn get_files_and_folders(path: _) -> _ {
    todo!()
}

fn get_file_sizes(file_and_folder_names: _) -> _ {
    todo!()
}

fn get_file_types(file_and_folder_names: _) -> _ {
    todo!()
}

fn get_files_and_folders_names(path: _) -> _ {
    todo!()
}
