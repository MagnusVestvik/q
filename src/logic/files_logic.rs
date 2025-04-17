use std::fs;
use std::fs::ReadDir;
fn read_files_in_dir(path: String) -> ReadDir {
    let files = fs::read_dir(path).unwrap();
    files
}

pub fn get_paths(path: String) -> Vec<String> {
    let dir_files = read_files_in_dir(path);
    dir_files
        .filter_map(|entry| {
            entry
                .ok()
                .and_then(|e| e.file_name().to_str().map(String::from))
        })
        .collect()
}

pub fn detect_path_type(path: String) -> String {
    match path.strip_suffix("/") {
        Some(folder) => "folder".to_string(),
        None => "file".to_string(),
    }
}

pub fn get_path_size(path: String) -> u64 {
    let metadata = fs::metadata(path).unwrap();
    metadata.len()
}

pub fn collect_path_outputs(path: String) -> Vec<Vec<String>> {
    get_paths(path)
        .iter()
        .map(|path_name| {
            let path_type = detect_path_type(path_name.clone());
            let path_size = get_path_size(path_name.clone());

            vec![path_name.clone(), path_type, path_size.to_string()]
        })
        .collect()
}
