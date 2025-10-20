#include <fstream>
#include <iostream>
#include <nlohmann/json.hpp>

using json = nlohmann::json;

int main() {
    //todo: remove hardcode filename
    std::ifstream input_file("./report.json");
    if (!input_file.is_open()) {
        std::cerr << "Error: Could not open file for reading.\n";
        return 1;
    }

    //import json contents to memory
    json data;
    input_file >> data;
    input_file.close();

    //extract relevant fields
    json cleaned_data = {
        {"file_count", data["file_count"]},
        {"language_stats", data["language_stats"]},
        {"total_lines", data["total_lines"]},
        {"code_lines", data["code_lines"]},
        {"comment_lines", data["comment_lines"]},
        {"blank_lines", data["blank_lines"]},
        {"commit_count", data["commit_count"]},
        {"contributors", data["contributors"]},
        {"last_activity", data["last_activity"]}
    };

    std::ofstream output_file("cleaned_report.json");
    if (!output_file.is_open()) {
        std::cerr << "Error: Could not open file for writing.\n";
        return 1;
    }

    output_file << std::setw(4) << cleaned_data << std::endl;
    output_file.close();

    std::cout << "Cleaned JSON saved to cleaned_report.json\n";
    return 0;
}
