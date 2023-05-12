
ExUnit.start

defmodule DayOne do
  def rlist([f | rest], total) do
    if f >= Enum.at(rest, 0) do
      rlist(rest, total)
    else
      rlist(rest, total + 1)
    end
  end

  def rlist([], total) do
    total - 1
  end
end

values = File.read!("./input")
|> String.split("\n", trim: true)
|> Enum.map(&String.to_integer/1)

IO.puts DayOne.rlist(values, 0)


defmodule DayOneTest do
  use ExUnit.Case

  test "normal" do
    assert DayOne.rlist([199, 200, 208, 210, 200, 207, 240, 269, 260, 263], 0) == 7
  end

  test "same" do
    assert DayOne.rlist([199, 200, 208, 208, 210, 200, 207, 240, 269, 260, 263], 0) == 7
  end

  test "smaller" do
    assert DayOne.rlist([1, 2, 3], 0) == 2
  end

  test "smaller same" do
    assert DayOne.rlist([1, 1, 2, 3], 0) == 2
  end

end
