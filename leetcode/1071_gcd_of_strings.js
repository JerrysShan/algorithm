var gcdOfStrings = function(str1, str2) {
    if(!str1||!str2){
        return '';
    }
    let str='';
    let temp='';
    let i=0;
    while(i<str1.length&&i<str2.length){
        if(str1[i]===str2[i]){
          temp+=str1[i];
            if(test(temp,str1)&&test(temp,str2)){
                if(temp.length>str.length){
                    str=temp;
                }
            }
            i++;
        }else{
            break;
        }
    }

    return str;
};

function test(sub,all){
    let temp=sub;
    while(temp.length<=all.length){
        if(temp==all){
            return true;
        }
        temp+=sub;
    }
    return false;
}

console.log(gcdOfStrings('ABCABC','ABC'));
console.log(gcdOfStrings('ABABAB','ABAB'));