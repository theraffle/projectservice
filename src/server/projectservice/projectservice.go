package projectservice

import (
	"context"
	"fmt"
	"github.com/theraffle/projectservice/src/database"
	"github.com/theraffle/projectservice/src/genproto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	log = logf.Log.WithName("user-service")
)

// ProjectService is a struct that have project service gRPC methods
type ProjectService struct{}

// CreateProject creates a project in project table and returns generated project ID
func (p ProjectService) CreateProject(ctx context.Context, request *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	db, err := database.Connect()
	if err != nil {
		log.Error(err, "database connection error")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error(err, "db closing error")
			os.Exit(1)
		}
	}()

	query := fmt.Sprintf("INSERT INTO project(name, chain_id, raffle_contract) VALUES('%s', %d, '%s')", request.ProjectName, request.ChainID, request.RaffleContract)
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Error(err, "create project error")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	n, _ := result.RowsAffected()
	if n == 1 {
		response := &pb.CreateProjectResponse{}
		query = fmt.Sprintf("SELECT id FROM project WHERE name = '%s', chain_id = %d, raffle_contract = '%s'", request.ProjectName, request.ChainID, request.RaffleContract)
		if err = db.QueryRowContext(ctx, query).Scan(&response.ProjectID); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		return response, nil
	}

	log.Error(err, "check database project table")
	return nil, status.Errorf(codes.Internal, "project insertion error: check database project table")
}

// GetProject returns project's info from project table
func (p ProjectService) GetProject(ctx context.Context, request *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	db, err := database.Connect()
	if err != nil {
		log.Error(err, "database connection error")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error(err, "db closing error")
			os.Exit(1)
		}
	}()

	query := fmt.Sprintf("SELECT * FROM project WHERE id = %d", request.ProjectID)
	row := db.QueryRowContext(ctx, query)

	response := &pb.GetProjectResponse{}
	err = row.Scan(&response.Project.ProjectID, &response.Project.ProjectName, &response.Project.ChainID, &response.Project.RaffleContract)

	if err != nil {
		log.Error(err, "getting project from db error")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return response, nil
}

// GetAllProjects returns all project list in project table
func (p ProjectService) GetAllProjects(ctx context.Context, empty *pb.Empty) (*pb.GetAllProjectResponse, error) {
	db, err := database.Connect()
	if err != nil {
		log.Error(err, "database connection error")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error(err, "db closing error")
			os.Exit(1)
		}
	}()

	query := "SELECT * FROM project"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Error(err, "get projects query error")
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	var projects []*pb.Project
	for rows.Next() {
		var project *pb.Project
		err = rows.Scan(&project.ProjectID, &project.ProjectName, &project.ChainID, &project.RaffleContract)
		if err != nil {
			log.Error(err, "get projects scan error")
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		projects = append(projects, project)
	}
	return &pb.GetAllProjectResponse{Projects: projects}, nil
}

// UpdateProject is method for updating a project
func (p ProjectService) UpdateProject(ctx context.Context, request *pb.UpdateProjectRequest) (*pb.GetProjectResponse, error) {
	// TODO implement later when project updating rule is defined

	return nil, nil
}
