// recording the time and counting the prime number till 1M (100,000,000)
// Total Number of prime numbers: 5761456
// total time taken 148037
const MIL = 100000000

function isPrime(num) {
    for (let i = 2; i * i <= num; i += 1) {
        if (num % i === 0) {
            return false;
        }
    }
    return true;
}

function main() {
    let totalPrimeNumbers = 0;
    const start = Date.now();
    let inBetweenTime = start; 
    for (let i = 1; i <= MIL; i += 1) {
        totalPrimeNumbers += isPrime(i) ? 1 : 0;
        
        // In this the output in ``time : 6425 -    0  -     664580 `` which means zero in the middler 
        // which mean that this is going to another thread and when console log runs that inBetweenTime is set to the current time 
        // by the next statement
        // i % 10000000 === 0 ? console.log(`time : ${Date.now() - start} -\t ${Date.now() - inBetweenTime}  -\t  ${totalPrimeNumbers}`) : '';
        // inBetweenTime = Date.now();

        // can be solved like this
        if (i % 10000000 === 0) {
            console.log(`time : ${Date.now() - start} -\t ${Date.now() - inBetweenTime}  -\t  ${totalPrimeNumbers}`);
            inBetweenTime = Date.now();
        }
    }
    console.log("Total Number of prime numbers : ", totalPrimeNumbers)
}

const start = Date.now();
main()
console.log(`total time taken ${Date.now() - start}ms`)