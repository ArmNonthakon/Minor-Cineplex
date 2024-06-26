import { memo, useEffect, useState } from 'react';
import './seats.scss';
import { checkArray } from '../../utils/filter';

interface Input {
    row: number;
    col: number;
    sendSeat : (seat:string[])=> void
}

const LoopRow = memo(({ row }: { row: number }) => {
    let result = [];
    for (let i = 0; i < row; i++) {
        let text: string = String.fromCharCode(65 + i);
        result.push(<h4 key={i}>{text}</h4>);
    }
    return <>{result}</>;
});


export const Seats = memo(({ row, col ,sendSeat}: Input) => {
    console.log('recomponent seats')
    const [seatReserving, setSeatReserving] = useState<string[]>([])
    const LoopSeat = ({ row, col }: { row: number, col: number }) => {
        let h = [];
        let already: string[] = [];
        for (let i = 0; i < row; i++) {
            let acii = 65 + i;
            let text: string = String.fromCharCode(acii);
            for (let j = 1; j <= col; j++) {
                let seatnumber: string = text + j;
                if (checkArray(seatnumber, already) && !checkArray(seatnumber, seatReserving)) {
                    h.push(
                        <img
                            className='seat-user'
                            key={seatnumber}
                            src="/user_seat.png"
                            width="31px"
                            alt=""
                        />
                    );
                }
                else if (checkArray(seatnumber, seatReserving)) {
                    h.push(
                        <img
                            className='seat-reserve'
                            key={seatnumber}
                            src="/check.png"
                            width="24px"
                            alt=""
                            onClick={() => {
                                setSeatReserving(oldValues => {
                                    let newData = oldValues.filter(data => data !== seatnumber)
                                    sendSeat(newData)
                                    return newData
                                })

                            }}
                        />
                    );
                }
                else {
                    h.push(
                        <img
                            key={seatnumber}
                            src="/sofa.png"
                            width="30px"
                            alt=""
                            onClick={() => {
                                sendSeat([...seatReserving,seatnumber])
                                setSeatReserving([...seatReserving, seatnumber])
                                
                            }}
                        />
                    );
                }
            }
        }
        return <>{h}</>;
    };
    useEffect(() => {
       
    }, []);

    return (
        <div className='section-seats'>
            <div className='section-seats-screen'>
                SCREEN
            </div>
            <div className='section-seats-seat'>
                <div className='seats-number'>
                    <LoopRow row={row} />
                </div>
                <div className='seats' style={{ gridTemplateColumns: `repeat(${col}, ${100 / col - 2}%)` }}>
                    <LoopSeat row={row} col={col} />
                </div>
            </div>
        </div>
    );
});
