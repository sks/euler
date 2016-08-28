package io.github.sks.euler.scala

/**
 * ans: 233168
 */
object Problem1 {
  
  def main(args: Array[String]): Unit = {
    println(solutionsForProblem1())
  }
  
  def isMultiple(x: Int): Boolean = 
    return (x%3) ==0 || (x%5==0)

  def solutionsForProblem1() = {
    var total = 0 
    for(num <-  1 to 999) {
      if (isMultiple(num)) {
        total+=num
      }
    }
    total
  }
}