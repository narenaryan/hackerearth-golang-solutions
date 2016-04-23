function factorial(n:number): number{
if(n == 1){
    return 1; 
  }
else{ 
    return n * factorial(n - 1);
  }
}

    console.log(factorial(5));