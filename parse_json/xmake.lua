add_requires("nlohmann_json")

add_rules("mode.debug", "mode.release")

target("parse_json")
    set_kind("binary")
    set_languages("c++20")
    add_files("src/*.cpp")
    add_packages("nlohmann_json")
