interface SideNavButtonProps {
  name: string;
  icon: React.ReactNode;
  func?: () => void;
}

interface SideNavGroupProps {
  groupName: string;
  options: SideNavButtonProps[];
}

interface SideNavProps {
  isSideNavOpen: boolean;
  groups: SideNavGroupProps[];
}

export default function SideNav({ isSideNavOpen, groups }: SideNavProps) {
  return (
    <div className="Dashboard-side-nav">
      {groups.map((group) => (
        <div key={group.groupName}>
          {isSideNavOpen && <h1>{group.groupName}</h1>}
          {group.options.map((option) => (
            <div
              className="Side-nav-button"
              key={option.name}
              onClick={option.func}
            >
              <i>{option.icon}</i>
              {isSideNavOpen && <p>{option.name}</p>}
            </div>
          ))}
        </div>
      ))}
    </div>
  );
}
