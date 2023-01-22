package battleships

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Board
	}{
		{
			name: "New 1x1 board",
			args: args{size: 1},
			want: &Board{
				size: 1,
				tiles: [][]Tile{
					{
						{
							x:     0,
							y:     0,
							state: EmptyState,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_calculateSingleMoveForShip(t *testing.T) {
	type fields struct {
		size  int
		tiles [][]Tile
		bombs []*Bomb
	}
	type args struct {
		ship *Ship
		move *Move
	}
	tests := []struct {
		name string
		args args
		want Move
	}{
		{
			name: "Len2 ship can move 1 to the front with right rotation",
			args: args{
				ship: &Ship{
					length:   2,
					rotation: RightRotation,
					gridPos: []ShipPosition{
						{
							x:       5,
							y:       5,
							isFront: true,
						},
						{
							x:       4,
							y:       5,
							isFront: false,
						},
					},
				},
				move: &Move{
					xOffset:    1,
					yOffset:    0,
					isPossible: false,
				},
			},
			want: Move{
				xOffset:    1,
				yOffset:    0,
				isPossible: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoard(10)
			if got := b.calculateSingleMoveForShip(tt.args.ship, tt.args.move); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSingleMoveForShip() = %v, want %v", got, tt.want)
			}
		})
	}
}
