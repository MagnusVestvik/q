use crate::domain::display::printer::Printer;
use crate::logic::path_handler::get_files_as_list;
use std::env;
mod domain;
mod logic;
fn main() {
    let args: Vec<String> = env::args().collect();
    let files = get_files_as_list(".".to_string());
    println!("{}, with args: {:?}", files.join(", "), args) // NB: the :? is for things that
    // implements the debug trait and
    // regular is for things that
    // inplements the Display trait
    /*
     * TODO: use pattern matching for command line arguments that fire the different functions
     */
}

/*
 * switch statements med args
 */
