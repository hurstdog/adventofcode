#!/usr/bin/env python3

def read_input_sections(filename: str) -> tuple[list[str], list[str]]:
    """
    Read a file containing two sections separated by a blank line.
    
    Args:
        filename: Path to the input file
        
    Returns:
        A tuple containing two lists of strings:
        - First list: lines from the first section (ordering_rules)
        - Second list: lines from the second section (page_listings)
        
    Raises:
        FileNotFoundError: If the input file doesn't exist
        ValueError: If the file doesn't contain exactly two sections
    """
    try:
        with open(filename, 'r') as f:
            content = f.read().strip()
            
        # Split into sections by blank lines
        sections = [section.strip() for section in content.split('\n\n')]
        
        if len(sections) != 2:
            raise ValueError(f"Expected 2 sections, found {len(sections)}")
            
        # Split each section into lines and remove empty lines
        ordering_rules = [line for line in sections[0].split('\n') if line]
        page_listings = [line for line in sections[1].split('\n') if line]
        
        return ordering_rules, page_listings
        
    except FileNotFoundError:
        raise FileNotFoundError(f"Input file '{filename}' not found")

def parse_rule(rule: str) -> tuple[str, str]:
    """Parse a rule into its two numbers."""
    return tuple(rule.split('|'))

def check_order(numbers: list[str], first: str, second: str) -> bool:
    """
    Check if first appears before second in the list of numbers.
    Returns True if:
    - first appears before second, or
    - either number is not in the list
    Returns False only if both numbers are present and in the wrong order.
    """
    try:
        first_idx = numbers.index(first)
        second_idx = numbers.index(second)
        # Only fail if both numbers are present and in wrong order
        return first_idx < second_idx
    except ValueError:
        # If either number is not in the list, the rule passes
        return True

def get_middle_number(listing: str) -> int:
    """
    Get the middle number from a comma-separated list of numbers.
    For lists with odd length, returns the middle number.
    """
    numbers = listing.split(',')
    middle_idx = len(numbers) // 2
    return int(numbers[middle_idx])

def validate_page_listing(listing: str, rules: list[str]) -> bool:
    """
    Check if a page listing satisfies all ordering rules.
    
    Args:
        listing: A comma-separated string of numbers
        rules: List of rules in the format "X|Y"
        
    Returns:
        True if the listing satisfies all rules, False otherwise
    """
    # Split the listing into a list of numbers
    numbers = listing.split(',')
    
    # Check each rule
    for rule in rules:
        first, second = parse_rule(rule)
        if not check_order(numbers, first, second):
            print(f"  Failed rule: {first} must come before {second} in {listing}")
            return False
    return True

def fix_page_listing(listing: str, rules: list[str]) -> str:
    """
    Fix an invalid page listing by only swapping numbers that violate the rules.
    Preserves the original order of numbers as much as possible.
    
    Args:
        listing: A comma-separated string of numbers
        rules: List of rules in the format "X|Y"
        
    Returns:
        A new comma-separated string with minimal changes to satisfy the rules
    """
    numbers = listing.split(',')
    fixed = False
    
    # Keep swapping numbers until all rules are satisfied
    while not fixed:
        fixed = True
        for rule in rules:
            first, second = parse_rule(rule)
            try:
                first_idx = numbers.index(first)
                second_idx = numbers.index(second)
                # If both numbers exist and are in wrong order, swap them
                if first_idx > second_idx:
                    numbers[first_idx], numbers[second_idx] = numbers[second_idx], numbers[first_idx]
                    fixed = False
            except ValueError:
                # If either number is not in the list, skip this rule
                continue
    
    return ','.join(numbers)

def main():
    try:
        # day 5 part 1 answer: 7074
        ordering_rules, page_listings = read_input_sections('input.txt')
        
        print("Page Listing Validation:")
        middle_sum = 0
        valid_count = 0
        fixed_middle_sum = 0
        invalid_count = 0
        
        for i, listing in enumerate(page_listings, 1):
            is_valid = validate_page_listing(listing, ordering_rules)
            status = "PASS" if is_valid else "FAIL"
            
            if is_valid:
                middle_num = get_middle_number(listing)
                middle_sum += middle_num
                valid_count += 1
                print(f"Listing {i}: {listing} -> {status} (middle: {middle_num})")
            else:
                invalid_count += 1
                print(f"Listing {i}: {listing} -> {status}")
                # Try to fix the invalid listing
                fixed_listing = fix_page_listing(listing, ordering_rules)
                if validate_page_listing(fixed_listing, ordering_rules):
                    fixed_middle_num = get_middle_number(fixed_listing)
                    fixed_middle_sum += fixed_middle_num
                    print(f"  Fixed: {fixed_listing} (middle: {fixed_middle_num})")
                else:
                    print("  Could not fix listing")
            
        print(f"\nFound {valid_count} valid sequences")
        print(f"Fixed {invalid_count} invalid sequences")
        print(f"Sum of original middle numbers: {middle_sum}")
        print(f"Sum of fixed middle numbers: {fixed_middle_sum}")
            
    except Exception as e:
        print(f"Error: {e}")
        return 1
    return 0

if __name__ == '__main__':
    exit(main()) 