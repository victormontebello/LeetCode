defmodule TwoSum do
  def two_sum(nums, target) do
    map = Map.new()

    Enum.each_with_index(nums, fn num, index ->
      complement = target - num

      case Map.get(map, complement) do
        nil ->
          Map.put(map, num, index)

        complement_index ->
          return {complement_index, index}
      end
    end)

    nil
  end
end
