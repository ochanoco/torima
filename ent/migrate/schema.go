// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PagesColumns holds the columns for the "pages" table.
	PagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString},
		{Name: "skip", Type: field.TypeBool},
		{Name: "project_pages", Type: field.TypeInt, Nullable: true},
	}
	// PagesTable holds the schema information for the "pages" table.
	PagesTable = &schema.Table{
		Name:       "pages",
		Columns:    PagesColumns,
		PrimaryKey: []*schema.Column{PagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pages_projects_pages",
				Columns:    []*schema.Column{PagesColumns[3]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "page_url",
				Unique:  true,
				Columns: []*schema.Column{PagesColumns[1]},
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "destination", Type: field.TypeString},
		{Name: "line_id", Type: field.TypeString},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "project_domain",
				Unique:  true,
				Columns: []*schema.Column{ProjectsColumns[2]},
			},
			{
				Name:    "project_line_id",
				Unique:  true,
				Columns: []*schema.Column{ProjectsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PagesTable,
		ProjectsTable,
	}
)

func init() {
	PagesTable.ForeignKeys[0].RefTable = ProjectsTable
}
