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
