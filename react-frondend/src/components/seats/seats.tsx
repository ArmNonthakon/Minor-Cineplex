import { useEffect } from 'react';
import './seats.scss';

interface Input {
    row: number;
    col: number;
}

const LoopRow = ({ row }: { row: number }) => {
    let result = [];
    for (let i = 0; i < row; i++) {
        let text: string = String.fromCharCode(65 + i);
        result.push(<h4 key={i}>{text}</h4>);
    }
    return <>{result}</>;
};

const LoopSeat = ({ row, col }: { row: number, col: number }) => {
    let h = [];
    let already = "";
    for (let i = 0; i < row; i++) {
        let acii = 65 + i;
        let text: string = String.fromCharCode(acii);
        for (let j = 1; j <= col; j++) {
            let seatnumber: string = text + j;
            if (seatnumber === already) {
                h.push(
                    <img
                        className='seat-user'
                        key={seatnumber}
                        src="/user_seat.png"
                        width="31px"
                        alt=""
                    />
                );
            } else {
                h.push(
                    <img
                        key={seatnumber}
                        src="/sofa.png"
                        width="30px"
                        alt=""
                    />
                );
            }
        }
    }
    return <>{h}</>;
};

export const Seats = ({ row, col }: Input) => {
    useEffect(() => {
        // You can add any side-effects or initialization code here if needed.
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
};
