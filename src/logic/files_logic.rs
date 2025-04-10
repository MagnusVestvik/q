use std::fs;
use std::fs::ReadDir;
fn read_files_in_dir(path: String) -> ReadDir {
    let files = fs::read_dir(path).unwrap();
    files
}

pub fn get_files_as_list(path: String) -> Vec<String> {
    let dir_files = read_files_in_dir(path);
    dir_files
        .filter_map(|entry| {
            entry
                .ok()
                .and_then(|e| e.file_name().to_str().map(String::from))
        })
        .collect()
}
