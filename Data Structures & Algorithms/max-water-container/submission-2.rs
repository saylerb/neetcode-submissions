impl Solution {
    pub fn max_area(heights: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut right = heights.len() - 1;
        let mut max_volume = 0;

        while left < right {
            let mut volume = 0;
            let width = right - left;
            let left_height = heights[left];
            let right_height = heights[right];

            if left_height > right_height {
                volume = width as i32 * right_height
            } else {
                volume = width as i32 * left_height
            }
            if volume > max_volume {
                max_volume = volume
            }
            if left_height < right_height {
                left += 1
            } else {
                right -= 1
            }
        } 
        return max_volume as i32
    }
}