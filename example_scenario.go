package main

type RGB struct{
  R uint8 
  G uint8 
  B uint8
}

type scaling_factor struct{
  Rf float32
  Gf float32
  Bf float32
}

func create_grayscale_matrix(size int){  
  //populate_image_matrix
  image := [][]RGB{
    {RGB{255,0,0},RGB{0,255,0},RGB{0,0,255}},
    {RGB{255,255,0},RGB{0,255,255},RGB{255,0,255}},
    {RGB{0,0,0},RGB{255,255,255},RGB{255,128,255}},
  } 
  grayscale := scaling_factor{0.299,0.5,0.432}
}
