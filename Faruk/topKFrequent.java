class Solution {
    public int[] topKFrequent(int[] nums, int k) {

        PriorityQueue<Map.Entry<Integer, Integer>> minHeap = new PriorityQueue<>((a, b) -> a.getValue() - b.getValue());

        Map<Integer, Integer> map = new HashMap<>();
        int[] res = new int[k];

        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {

            if (minHeap.size() < k) {
                minHeap.add(entry);
            } else if (minHeap.peek().getValue() < entry.getValue()) {

                minHeap.poll();
                minHeap.add(entry);

            }
        }

        for (int i = 0; i < k; i++) {
            res[i] = minHeap.poll().getKey();
        }

        return res;

    }
}