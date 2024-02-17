package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		/**
		 * ここでテストケースを定義する
		 * テストケースは、name, args, wantの3つの要素を持つ構造体のスライスである
		 * name: テストケースの名前
		 * args: テストケースの引数
		 * want: テストケースの期待値
		 */
		{
			name: "1+2=3",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "3+5=8",
			args: args{x: 3, y: 5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1/2=0.5",
			args: args{x: 1, y: 2},
			want: 0.5,
		},
		{
			name: "2/0=0",
			args: args{x: 2, y: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
