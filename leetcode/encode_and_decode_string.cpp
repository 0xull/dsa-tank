#include <string>
#include <vector>

// Leetcode 271 Encode and Decode Strings
class Solution {
public:
  std::string encode(std::vector<std::string> input) {
    std::string output;

    for (std::string str : input) {
      output += std::to_string(str.size()) + "#" + str;
    }

    return output;
  }

  std::vector<std::string> decode(std::string encoded) {
    int i = 0;
    std::vector<std::string> output;

    while (i < encoded.size()) {
      int j = i;
      while (encoded[j] != '#')
        j++;

      int substr_len = std::stoi(encoded.substr(i, j - i));
      output.push_back(encoded.substr(j + 1, substr_len));
      i = (j + 1 + substr_len);
    }

    return output;
  }
};
