

export const RemoveDuplicate = (data: string[]) =>{
    return data.filter((value,index)=> data.indexOf(value)===index);
}

export const checkArray = (dataCheck: string , array: string[])=>{
    let result = false;
    array.map((e)=>{
        if(e == dataCheck){
            result = true
        }
    })
    return result
}