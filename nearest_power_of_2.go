package main

func nearest_power_of_2(n int) int{
  pow := 2
  for i:=2;i<n;i*=2 {
    pow = i*2
  }
  return pow
}
