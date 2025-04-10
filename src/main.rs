use crate::logic::files_logic::get_files_as_list;
mod logic;
fn main() {
    let files = get_files_as_list(".".to_string());
    println!("{}", files.join(", "))
}
