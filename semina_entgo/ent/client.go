// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"semina_entgo/ent/migrate"

	"semina_entgo/ent/car"
	"semina_entgo/ent/card"
	"semina_entgo/ent/group"
	"semina_entgo/ent/pet"
	"semina_entgo/ent/tester"
	"semina_entgo/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// Tester is the client for interacting with the Tester builders.
	Tester *TesterClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Car = NewCarClient(c.config)
	c.Card = NewCardClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Pet = NewPetClient(c.config)
	c.Tester = NewTesterClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Car:    NewCarClient(cfg),
		Card:   NewCardClient(cfg),
		Group:  NewGroupClient(cfg),
		Pet:    NewPetClient(cfg),
		Tester: NewTesterClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Car:    NewCarClient(cfg),
		Card:   NewCardClient(cfg),
		Group:  NewGroupClient(cfg),
		Pet:    NewPetClient(cfg),
		Tester: NewTesterClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Car.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the dto clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	for _, n := range []interface{ Use(...Hook) }{
		c.Car, c.Card, c.Group, c.Pet, c.Tester, c.User,
	} {
		n.Use(hooks...)
	}
}

// Intercept adds the query interceptors to all the dto clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	for _, n := range []interface{ Intercept(...Interceptor) }{
		c.Car, c.Card, c.Group, c.Pet, c.Tester, c.User,
	} {
		n.Intercept(interceptors...)
	}
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CarMutation:
		return c.Car.mutate(ctx, m)
	case *CardMutation:
		return c.Card.mutate(ctx, m)
	case *GroupMutation:
		return c.Group.mutate(ctx, m)
	case *PetMutation:
		return c.Pet.mutate(ctx, m)
	case *TesterMutation:
		return c.Tester.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `car.Intercept(f(g(h())))`.
func (c *CarClient) Intercept(interceptors ...Interceptor) {
	c.inters.Car = append(c.inters.Car, interceptors...)
}

// Create returns a builder for creating a Car dto.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CarClient) MapCreateBulk(slice any, setFunc func(*CarCreate, int)) *CarCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CarCreateBulk{err: fmt.Errorf("calling to CarClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CarCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCar},
		inters: c.Interceptors(),
	}
}

// Get returns a Car dto by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Car.
func (c *CarClient) QueryOwner(ca *Car) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(car.Table, car.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, car.OwnerTable, car.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// Interceptors returns the client interceptors.
func (c *CarClient) Interceptors() []Interceptor {
	return c.inters.Car
}

func (c *CarClient) mutate(ctx context.Context, m *CarMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CarCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CarUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CarDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Car mutation op: %q", m.Op())
	}
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `card.Hooks(f(g(h())))`.
func (c *CardClient) Use(hooks ...Hook) {
	c.hooks.Card = append(c.hooks.Card, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `card.Intercept(f(g(h())))`.
func (c *CardClient) Intercept(interceptors ...Interceptor) {
	c.inters.Card = append(c.inters.Card, interceptors...)
}

// Create returns a builder for creating a Card dto.
func (c *CardClient) Create() *CardCreate {
	mutation := newCardMutation(c.config, OpCreate)
	return &CardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Card entities.
func (c *CardClient) CreateBulk(builders ...*CardCreate) *CardCreateBulk {
	return &CardCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CardClient) MapCreateBulk(slice any, setFunc func(*CardCreate, int)) *CardCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CardCreateBulk{err: fmt.Errorf("calling to CardClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CardCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	mutation := newCardMutation(c.config, OpUpdate)
	return &CardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCard(ca))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id int) *CardUpdateOne {
	mutation := newCardMutation(c.config, OpUpdateOne, withCardID(id))
	return &CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	mutation := newCardMutation(c.config, OpDelete)
	return &CardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *CardClient) DeleteOneID(id int) *CardDeleteOne {
	builder := c.Delete().Where(card.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CardDeleteOne{builder}
}

// Query returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCard},
		inters: c.Interceptors(),
	}
}

// Get returns a Card dto by its id.
func (c *CardClient) Get(ctx context.Context, id int) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id int) *Card {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Card.
func (c *CardClient) QueryOwner(ca *Card) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(card.Table, card.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, card.OwnerTable, card.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CardClient) Hooks() []Hook {
	return c.hooks.Card
}

// Interceptors returns the client interceptors.
func (c *CardClient) Interceptors() []Interceptor {
	return c.inters.Card
}

func (c *CardClient) mutate(ctx context.Context, m *CardMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CardCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CardUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CardDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Card mutation op: %q", m.Op())
	}
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `group.Intercept(f(g(h())))`.
func (c *GroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.Group = append(c.inters.Group, interceptors...)
}

// Create returns a builder for creating a Group dto.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *GroupClient) MapCreateBulk(slice any, setFunc func(*GroupCreate, int)) *GroupCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &GroupCreateBulk{err: fmt.Errorf("calling to GroupClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*GroupCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a Group dto by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// Interceptors returns the client interceptors.
func (c *GroupClient) Interceptors() []Interceptor {
	return c.inters.Group
}

func (c *GroupClient) mutate(ctx context.Context, m *GroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Group mutation op: %q", m.Op())
	}
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `pet.Intercept(f(g(h())))`.
func (c *PetClient) Intercept(interceptors ...Interceptor) {
	c.inters.Pet = append(c.inters.Pet, interceptors...)
}

// Create returns a builder for creating a Pet dto.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PetClient) MapCreateBulk(slice any, setFunc func(*PetCreate, int)) *PetCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PetCreateBulk{err: fmt.Errorf("calling to PetClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PetCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id int) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *PetClient) DeleteOneID(id int) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePet},
		inters: c.Interceptors(),
	}
}

// Get returns a Pet dto by its id.
func (c *PetClient) Get(ctx context.Context, id int) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id int) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pet.OwnerTable, pet.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
}

// Interceptors returns the client interceptors.
func (c *PetClient) Interceptors() []Interceptor {
	return c.inters.Pet
}

func (c *PetClient) mutate(ctx context.Context, m *PetMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PetCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PetUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PetDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Pet mutation op: %q", m.Op())
	}
}

// TesterClient is a client for the Tester schema.
type TesterClient struct {
	config
}

// NewTesterClient returns a client for the Tester from the given config.
func NewTesterClient(c config) *TesterClient {
	return &TesterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tester.Hooks(f(g(h())))`.
func (c *TesterClient) Use(hooks ...Hook) {
	c.hooks.Tester = append(c.hooks.Tester, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tester.Intercept(f(g(h())))`.
func (c *TesterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Tester = append(c.inters.Tester, interceptors...)
}

// Create returns a builder for creating a Tester dto.
func (c *TesterClient) Create() *TesterCreate {
	mutation := newTesterMutation(c.config, OpCreate)
	return &TesterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tester entities.
func (c *TesterClient) CreateBulk(builders ...*TesterCreate) *TesterCreateBulk {
	return &TesterCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TesterClient) MapCreateBulk(slice any, setFunc func(*TesterCreate, int)) *TesterCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TesterCreateBulk{err: fmt.Errorf("calling to TesterClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TesterCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TesterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tester.
func (c *TesterClient) Update() *TesterUpdate {
	mutation := newTesterMutation(c.config, OpUpdate)
	return &TesterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *TesterClient) UpdateOne(t *Tester) *TesterUpdateOne {
	mutation := newTesterMutation(c.config, OpUpdateOne, withTester(t))
	return &TesterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TesterClient) UpdateOneID(id int) *TesterUpdateOne {
	mutation := newTesterMutation(c.config, OpUpdateOne, withTesterID(id))
	return &TesterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tester.
func (c *TesterClient) Delete() *TesterDelete {
	mutation := newTesterMutation(c.config, OpDelete)
	return &TesterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *TesterClient) DeleteOne(t *Tester) *TesterDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *TesterClient) DeleteOneID(id int) *TesterDeleteOne {
	builder := c.Delete().Where(tester.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TesterDeleteOne{builder}
}

// Query returns a query builder for Tester.
func (c *TesterClient) Query() *TesterQuery {
	return &TesterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTester},
		inters: c.Interceptors(),
	}
}

// Get returns a Tester dto by its id.
func (c *TesterClient) Get(ctx context.Context, id int) (*Tester, error) {
	return c.Query().Where(tester.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TesterClient) GetX(ctx context.Context, id int) *Tester {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TesterClient) Hooks() []Hook {
	return c.hooks.Tester
}

// Interceptors returns the client interceptors.
func (c *TesterClient) Interceptors() []Interceptor {
	return c.inters.Tester
}

func (c *TesterClient) mutate(ctx context.Context, m *TesterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TesterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TesterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TesterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TesterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Tester mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User dto.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given dto.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given dto.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given dto by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User dto by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCars queries the cars edge of a User.
func (c *UserClient) QueryCars(u *User) *CarQuery {
	query := (&CarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(car.Table, car.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CarsTable, user.CarsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPets queries the pets edge of a User.
func (c *UserClient) QueryPets(u *User) *PetQuery {
	query := (&PetClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PetsTable, user.PetsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := (&GroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.GroupsTable, user.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCard queries the card edge of a User.
func (c *UserClient) QueryCard(u *User) *CardQuery {
	query := (&CardClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(card.Table, card.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, user.CardTable, user.CardColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Car, Card, Group, Pet, Tester, User []ent.Hook
	}
	inters struct {
		Car, Card, Group, Pet, Tester, User []ent.Interceptor
	}
)
