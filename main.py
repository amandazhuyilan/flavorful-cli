import click

# Sample code snippet to demonstrate use of click
@click.command()
@click.argument('name')
@click.option('--greeting', '-g')
def main(name, greeting):
    click.echo("{}, {}!".format(greeting, name))

if __name__ == "__main__":
    main()