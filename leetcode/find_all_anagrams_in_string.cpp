#include <array>
#include <string>
#include <vector>

// O(1) optimization to the LC. 438 question (Find All Anagrams in a String)
class Solution {
public:
  std::vector<int> findAnagrams(std::string s, std::string p) {
      std::vector<int> result;

      if (s.size() < p.size()) return result;

      std::array<int, 26> freq_p = {0};
      std::array<int, 26> freq_w = {0};
      int matches = 0;
      
      for (char c : p) freq_p[c - 'a']++;
      for (int i = 0; i < 26; ++i) if (p[i] == 0) matches++;

      int j = 0; // tracks outgoing left character from the window slide;
      
      for (int i = 0; i < s.size(); ++i) {
          int in_char = s[i] - 'a';
          freq_w[in_char]++;

          if (freq_w[in_char] == freq_p[in_char]) matches++;
          else if (freq_w[in_char] == freq_p[in_char] + 1) matches--;

          if (i - j + 1 > p.size()) {
              int out_char = s[j] - 'a';
              freq_w[out_char]--;

              if (freq_w[out_char] == freq_p[out_char]) matches++;
              else if (freq_w[out_char] == freq_p[out_char] - 1) matches--;
          }

          if (matches == 26) result.push_back(j);
      }
      
      return result;
  }
}
