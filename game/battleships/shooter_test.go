package battleships

import (
	"testing"
)

func TestShooter_GetNewBomb(t *testing.T) {
	type fields struct {
		bombsDropped []Bomb
	}
	type args struct {
		board *Board
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pick random spot",
			fields: fields{
				bombsDropped: []Bomb{},
			},
			args: args{
				board: NewBoard(10),
			},
		},
		{
			name: "Previous bomb in top row",
			fields: fields{
				bombsDropped: []Bomb{
					{
						x:      0,
						y:      5,
						didHit: true,
					},
				},
			},
			args: args{
				board: NewBoard(10),
			},
		},
		{
			name: "Previous bomb in bottom row",
			fields: fields{
				bombsDropped: []Bomb{
					{
						x:      9,
						y:      5,
						didHit: true,
					},
				},
			},
			args: args{
				board: NewBoard(10),
			},
		},
		{
			name: "Previous bomb on the right",
			fields: fields{
				bombsDropped: []Bomb{
					{
						x:      5,
						y:      9,
						didHit: true,
					},
				},
			},
			args: args{
				board: NewBoard(10),
			},
		},
		{
			name: "Previous bomb on the left",
			fields: fields{
				bombsDropped: []Bomb{
					{
						x:      5,
						y:      0,
						didHit: true,
					},
				},
			},
			args: args{
				board: NewBoard(10),
			},
		},
		{
			name: "Previous bomb in the corner",
			fields: fields{
				bombsDropped: []Bomb{
					{
						x:      0,
						y:      0,
						didHit: true,
					},
				},
			},
			args: args{
				board: NewBoard(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shooter{
				bombsDropped: tt.fields.bombsDropped,
			}
			if got := s.GetNewBomb(tt.args.board); got.x < 0 || got.x > tt.args.board.size || got.y < 0 || got.y > tt.args.board.size || got.didHit != false {
				t.Errorf("GetNewBomb() = %v, invalid value", got)
			}
		})
	}
}
