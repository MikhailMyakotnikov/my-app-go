using OpenQA.Selenium;
using OpenQA.Selenium.Chrome;
using OpenQA.Selenium.Remote;
using OpenQA.Selenium.Support.UI;

namespace UiTests;

public class TeacherTests
{
    private IWebDriver driver;
    private WebDriverWait wait;
    private static readonly string? SeleniumServerUri = Environment
        .GetEnvironmentVariable("SELENIUM_SERVER_URI");
    private static readonly string? BaseAppUrl = Environment
        .GetEnvironmentVariable("APP_URL");
    private readonly List<string> createdTeachers = new();
    private const string CreateTeacherButtonTestId = "create-teacher-btn";
    private const string EditTeacherButtonTestId = "edit-teacher-action-btn";
    private const string ConfirmEditTeacherButtonTestId = "confirm-edit-teacher-btn";
    private const string DeleteTeacherButtonTestId = "delete-teacher-action-btn";
    [SetUp]
    public void SetUp()
    {
        if(SeleniumServerUri == null) throw new NullReferenceException(
            $"Ошибка: переменная {nameof(SeleniumServerUri)} не может хранить null");
        if(BaseAppUrl == null) throw new NullReferenceException(
            $"Ошибка: переменная {nameof(BaseAppUrl)} не может хранить null");
        
        Uri uri = new Uri(SeleniumServerUri);

        ChromeOptions options = new ChromeOptions();

        driver = new RemoteWebDriver(uri, options);

        wait = new WebDriverWait(driver, TimeSpan.FromSeconds(10));
    }
    [TearDown]
    public void TearDown()
    {
        foreach (var name in createdTeachers)
        {
            DeleteTeacherByName(name);
        }
        driver.Quit();
    }
    private IWebElement FindTeacherData(string name)
    {
        return wait.Until(driver => driver.FindElement(By.XPath(
            $"//td[text()='{name}']")));
    }
    private IWebElement FindTeacherRow(string name)
    {
        return wait.Until(driver => driver.FindElement(
            By.XPath($"//tr[td[text()='{name}']]")));
    }
    private void FindAndClickButton(string htmlTagAttribute)
    {
        IWebElement button = wait.Until(driver => driver.FindElement(By.CssSelector(
            $"[data-testid='{htmlTagAttribute}']")));

        button.Click();
    }
    private void FindAndClickButton(string htmlTagAttribute, IWebElement tableRow)
    {
        IWebElement button = wait.Until(_ => tableRow.FindElement
            (By.CssSelector($"[data-testid='{htmlTagAttribute}']")));

        button.Click();
    }
    private string CreateTeacher()
    {
        string appUrl = $"{BaseAppUrl}/teachers/create";
        driver.Navigate().GoToUrl(appUrl);
        wait.Until(driver => driver.Url.Contains("/teachers/create"));

        Console.WriteLine(driver.Url);
        Console.WriteLine(driver.PageSource);

        IWebElement nameInput = wait.Until(driver => driver.FindElement(
            By.Name("name")));

        string name = $"Teacher_{Guid.NewGuid().ToString().Substring(0, 8)}";
        //Example: Teacher_f3a1c5a4
        nameInput.SendKeys(name);

        FindAndClickButton(CreateTeacherButtonTestId);

        createdTeachers.Add(name);

        return name;
    }
    private void DeleteTeacherByName(string name)
    {
        driver.Navigate().GoToUrl($"{BaseAppUrl}/teachers");

        var rows = driver.FindElements(By.XPath($"//tr[td[text()='{name}']]"));

        if (rows.Count == 0) return;

        IWebElement row = rows[0];

        IWebElement button = row.FindElement
            (By.CssSelector($"[data-testid='{DeleteTeacherButtonTestId}']"));

        button.Click();

        IAlert alert = wait.Until(driver => driver.SwitchTo().Alert());
        alert.Accept();

        wait.Until(driver => driver.FindElements(By.XPath($"//td[text()='{name}']"))
            .Count == 0);
    }
    [Test]
    public void CreateTeacher_ShouldDisplayTeacherInTable()
    {
        string name = CreateTeacher();

        IWebElement createdTeacher = FindTeacherData(name);

        Assert.That(createdTeacher.Displayed, Is.True);
    }
    [Test]
    public void EditTeacher_ShouldUpdateTeacherName()
    {
        string name = CreateTeacher();

        IWebElement createdTeacher = FindTeacherData(name);

        Assert.That(createdTeacher.Displayed, Is.True);

        IWebElement row = FindTeacherRow(name);

        FindAndClickButton(EditTeacherButtonTestId, row);

        IWebElement nameInput = wait.Until(driver => driver.FindElement(
            By.Name("name")));

        string editedName = $"Edited_Teacher_{Guid.NewGuid().ToString().
            Substring(0, 8)}";//Example: Edited_Teacher_f3a1c5a4
        nameInput.Clear();
        nameInput.SendKeys(editedName);

        createdTeachers.Remove(name);
        createdTeachers.Add(editedName);

        FindAndClickButton(ConfirmEditTeacherButtonTestId);

        IWebElement editedTeacher = FindTeacherData(editedName);

        Assert.That(editedTeacher.Displayed, Is.True);
    }
    [Test]
    public void DeleteTeacher_ShouldRemoveTeacherFromTable()
    {
        string name = CreateTeacher();

        IWebElement createdTeacher = FindTeacherData(name);

        Assert.That(createdTeacher.Displayed, Is.True);

        IWebElement row = FindTeacherRow(name);

        FindAndClickButton(DeleteTeacherButtonTestId, row);

        IAlert alert = wait.Until(driver => driver.SwitchTo().Alert());
        alert.Accept();

        bool isNameDeleted = wait.Until(driver => driver.FindElements(
            By.XPath($"//td[text()='{name}']")).Count == 0);

        Assert.That(isNameDeleted, Is.True);
    }
}