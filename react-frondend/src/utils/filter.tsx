

export const RemoveDuplicate = (data: string[]) =>{
    return data.filter((value,index)=> data.indexOf(value)===index);
}