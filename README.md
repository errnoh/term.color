color
=====

image/color style model for 256 color terminals

To convert color.Color to closest 256 color terminal value you can do:

    // var c color.Color
    val := color.Term256Model.Convert(c).(color.Term256).Val

or you can get closest greyscale value with

    // var c color.Color
    val := color.Term256GreyscaleModel.Convert(c).(color.Term256).Val
