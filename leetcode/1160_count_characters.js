var countCharacters = function(words, chars) {
    let count=0;
    for(let i=0;i<words.length;i++){
        let item=words[i];
        let j=0;
        let temp=chars;
        for(;j<item.length;j++){
            if(temp.indexOf(item[j])>-1){
              let index=temp.indexOf(item[j]);
              temp=temp.slice(0,index)+temp.slice(index+1)
            }else{
                break;
            }
        }
        if(j===item.length){
            count+=j;
        }
    }
    return count;
};

console.log(countCharacters(["cat","bt","hat","tree"],'atach'))
 console.log(countCharacters(["hello","world","leetcode"],'welldonehoneyr'))