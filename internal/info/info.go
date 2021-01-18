package info

import (
	"context"
	"fmt"
	"time"

	grpc "github.com/antonioalfa22/egida/proto"
)

func GetWorkerInfo(hostslist []string, services string, packages string, hardening string) {
	if services != "" {
		if services == "stopped" {
			GetStoppedServices(hostslist)
		} else if services == "running" {
			GetRunningServices(hostslist)
		} else {
			GetAllServices(hostslist)
		}
	}
	if packages != "" {
		if packages == "all" {
			GetAllPackages(hostslist)
		}
	}
	if hardening != "" {
		if hardening == "lynis" {
			GetLynisScore(hostslist)
		}
	}
}

func GetAllServices(hosts []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, h := range hosts {
		client := CreateServicesClient(h)
		res, err := client.ListAllServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			fmt.Println("Error: can not connect to host ", h, " on port 8128")
		} else {
			fmt.Println("----> Host: ", h)
			for _, s := range res.Services {
				fmt.Println(s.Name, " [", s.Status, "]")
			}
		}
	}
}

func GetRunningServices(hosts []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, h := range hosts {
		client := CreateServicesClient(h)
		res, err := client.ListRunningServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			fmt.Println("Error: can not connect to host ", h, " on port 8128")
		} else {
			fmt.Println("----> Host: ", h)
			for _, s := range res.Services {
				fmt.Println(s.Name, " [", s.Status, "]")
			}
		}
	}
}

func GetStoppedServices(hosts []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, h := range hosts {
		client := CreateServicesClient(h)
		res, err := client.ListStoppedServices(ctx, &grpc.ListServicesReq{})
		if err != nil {
			fmt.Println("Error: can not connect to host ", h, " on port 8128")
		} else {
			fmt.Println("----> Host: ", h)
			for _, s := range res.Services {
				fmt.Println(s.Name, " [", s.Status, "]")
			}
		}
	}
}

func GetAllPackages(hosts []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, h := range hosts {
		client := CreatePackagesClient(h)
		res, err := client.ListAllPackages(ctx, &grpc.ListPackagesReq{})
		if err != nil {
			fmt.Println("Error: can not connect to host ", h, " on port 8128")
		} else {
			fmt.Println("----> Host: ", h)
			for _, s := range res.Packages {
				fmt.Println(s.Name)
			}
		}
	}
}

func GetLynisScore(hosts []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	fmt.Println("Getting Lynis scores ...")
	defer cancel()
	for _, h := range hosts {
		client := CreateHardeningClient(h)
		res, err := client.GetLynisScore(ctx, &grpc.ScoreReq{})
		if err != nil {
			fmt.Println("Error: can not connect to host ", h, " on port 8128")
		} else {
			fmt.Println("----> Host: ", h)
			fmt.Println("Score: ", res.Score)
			for _, l := range res.Log {
				fmt.Println(l)
			}
		}
	}
}
