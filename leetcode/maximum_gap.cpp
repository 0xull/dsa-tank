#include <algorithm>
#include <climits>
#include <vector>

class Solution {
  struct Bucket {
    int min_val = INT_MAX;
    int max_val = INT_MIN;
    bool used = false;
  };

public:
  int maximumGap(std::vector<int> &nums) {
    int n = nums.size();

    if (n < 2)
      return 0;

    auto [min_it, max_it] = std::minmax_element(nums.begin(), nums.end());
    int min_val = *min_it;
    int max_val = *max_it;

    if (min_val == max_val)
      return 0;

    long long range = static_cast<long long>(max_val) - min_val;

    long long bucket_width = std::max(1LL, (range + n - 2) / (n - 1));
    long long num_buckets = (range / bucket_width) + 1;

    std::vector<Bucket> buckets(num_buckets);

    for (int num : nums) {
      long long bucket_idx = (num - min_val) / bucket_width;
      buckets[bucket_idx].used = true;
      buckets[bucket_idx].min_val = std::min(buckets[bucket_idx].min_val, num);
      buckets[bucket_idx].max_val = std::max(buckets[bucket_idx].max_val, num);
    }

    int max_gap = 0;
    int previous_max = min_val;

    for (const auto& bucket : buckets) {
        if (!bucket.used) continue;

        max_gap = std::max(max_gap, bucket.min_val - previous_max);
        previous_max = bucket.max_val;
    }

    return max_gap;
  }
};
